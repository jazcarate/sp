// Package sourcing contains state and ways to change it
package sourcing

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	beginning = iota
	title
	glance
	participants
	details
	operations
	log
)

var (
	errNoHeader              = errors.New("the report did not start with the appropriate header")
	errNoName                = errors.New("the report did not have a name")
	errNoGlance              = errors.New("the report did not have an at-a-glance secction")
	errNoParticipants        = errors.New("the report did not have a participants secction")
	errNotAnchor             = errors.New("this is not a valid anchor")
	errNoDetails             = errors.New("the report did not have a details secction")
	errUnorderedParticipants = errors.New("the participant detail order did not match up")
	errWrongDetail           = errors.New("this is not a valid detail (comment)")
	errConfig                = errors.New("the operation configuration is wrong")
	errNoLastOp              = errors.New("there is no last operation link")
	errNotRecognized         = errors.New("operation not recognized")
)

// Parse an input onto a valir state.
func Parse(r io.Reader) (*State, error) {
	s := NewState()
	lineNumber := 0
	context := beginning

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		switch context {
		case beginning:
			if line != header {
				return s, errNoHeader
			}

			context = title
		case title:
			r := regexp.MustCompile(`^# (.+)$`)
			name := r.FindStringSubmatch(line)

			if len(name) != 2 {
				return s, errNoName
			}

			s.Name = name[1]

			scanner.Scan()

			context = glance
		case glance:
			if line == emptyHelp {
				context = operations

				scanner.Scan() // Operation header

				break
			}

			if line != glanceTitle {
				return s, errNoGlance
			}

			scanner.Scan() // at-a-glance-header
			scanner.Scan() // at-a-glance table hbody sepataror

			cols := strings.Split(scanner.Text(), "|")

			s.Balance = s.Balance.IncrD(len(cols) - 3)

			for y := 0; scanner.Scan(); y++ {
				glanceLine := scanner.Text()

				if glanceLine == "" {
					break
				}

				vals := strings.Split(glanceLine, "|")

				for x, cell := range vals[2 : len(vals)-1] {
					if x > y {
						if cell == "" {
							continue
						}

						v, err := strconv.Atoi(cell[1 : len(cell)-3])
						if err != nil {
							return s, fmt.Errorf("parsing glance value <%s> at #%d: %w", cell, lineNumber, err)
						}

						if cell[len(cell)-3:] == "↗" {
							v = -v
						}

						err = s.Balance.Set(x, y, v)
						if err != nil {
							return s, fmt.Errorf("could not set at a glance: %w", err)
						}
					}
				}
			}

			context = participants
		case participants:
			if line != participantsTitle {
				return s, errNoParticipants
			}

			scanner.Scan() // participants table header
			scanner.Scan() // participants table hbody sepataror

			for y := 0; scanner.Scan(); y++ {
				participantLine := scanner.Text()

				if participantLine == "" {
					break
				}

				vals := strings.Split(participantLine, "|")

				name, err := fromAnchor(vals[1])
				if err != nil {
					return s, fmt.Errorf("could not parse the participant at %d: %w", lineNumber, err)
				}

				split, err := strconv.Atoi(vals[2])
				if err != nil {
					return s, fmt.Errorf("parsing split value <%s> at #%d: %w", vals[2], lineNumber, err)
				}

				splitPercString := vals[3]

				splitPerc, err := strconv.Atoi(splitPercString[:len(splitPercString)-1])
				if err != nil {
					return s, fmt.Errorf("parsing split percentage value <%s> at #%d: %w", splitPercString, lineNumber, err)
				}

				newParcicipant := Participant{
					Name:            name,
					PublicKey:       "", // To be filled lated
					Split:           split,
					SplitPercentage: splitPerc,
				}

				s.Participants = append(s.Participants, newParcicipant)
			}

			context = details
		case details:
			if line != detailsTitle {
				return s, errNoDetails
			}

			scanner.Scan()

			for i := 0; i < len(s.Participants); i++ {
				participantName := scanner.Text()[4:]

				if participantName != s.Participants[i].Name {
					return s, errUnorderedParticipants
				}

				scanner.Scan()

				publicKey, err := removeComment("Public key", scanner.Text())
				if err != nil {
					return s, fmt.Errorf("coudn't get the public key for <%s>: %w", participantName, err)
				}

				s.Participants[i].PublicKey = publicKey

				for scanner.Scan() { // ignore everything else
					l := scanner.Text()
					if (len(l) > 4 && l[:4] == "### ") || l == operationsTitle {
						break
					}
				}
			}

			context = operations
		case operations:
			config := strings.Split(line, "**")
			if len(config) != 3 {
				return s, errConfig
			}

			s.Configuration = config[1]

			scanner.Scan()

			context = log
		case log:
			if line == "### Log" {
				return s, nil
			}

			r := regexp.MustCompile(`^### Log \[\(go to the last ⬇\)\]\(#op-(\d+)\)$`)
			res := r.FindStringSubmatch(line)

			lastOp, err := strconv.Atoi(res[1])
			if err != nil || len(res) != 2 {
				return nil, errNoLastOp
			}

			s.LastOp = lastOp

			scanner.Scan() // log table header
			scanner.Scan() // log table hbody sepataror

			for scanner.Scan() {
				vals := strings.Split(scanner.Text(), "|")

				idS, err := fromAnchor(vals[1])
				if err != nil {
					return s, fmt.Errorf("parsing log id anchor value <%s>: %w", vals[2], err)
				}

				id, err := strconv.Atoi(idS)
				if err != nil {
					return s, fmt.Errorf("parsing log id value <%s>: %w", idS, err)
				}

				on, err := time.Parse(time.RFC3339, vals[3])
				if err != nil {
					return s, fmt.Errorf("parsing log id #%d time: %w", id, err)
				}

				op, by, signature, err := parseOperation(vals[2])
				if err != nil {
					return s, fmt.Errorf("parsing operation id #%d: %w", id, err)
				}

				valid := false
				if vals[5] == "✅" {
					valid = true
				}

				newLog := LogEvent{
					ID:        id,
					By:        by,
					Operation: op,
					On:        on.Unix(),
					Note:      strings.ReplaceAll(vals[4], "<br />", "\n"),
					Signature: signature,
					Valid:     valid,
				}

				s.Log = append(s.Log, newLog)
			}
		}
	}

	err := scanner.Err()
	if err != nil {
		return s, fmt.Errorf("parsing line %d: %w", lineNumber, err)
	}

	return s, nil
}

func fromAnchor(input string) (string, error) {
	r := regexp.MustCompile(`^\[(.+)\]\(#.+\)`)
	res := r.FindStringSubmatch(input)

	if res == nil || len(res) != 2 {
		return "", fmt.Errorf("<%s> is not a valid anchor: %w", input, errNotAnchor)
	}

	return res[1], nil
}

func removeComment(key, input string) (string, error) {
	r := regexp.MustCompile(`^<!-- ` + key + ` ([a-zA-Z/0-9+]+) -->$`)
	res := r.FindStringSubmatch(input)

	if res == nil || len(res) != 2 {
		return "", fmt.Errorf("<%s> is not a valid %s: %w", input, key, errWrongDetail)
	}

	return res[1], nil
}

func parseOperation(input string) (StateChanger, string, string, error) {
	opLogS := strings.Split(input, " <!-- Sign ")
	signMeta := strings.Split(opLogS[1], " ")

	by := signMeta[0]
	signature := signMeta[1]

	opsS := strings.Split(opLogS[0], "<br />")

	var ops []StateChanger = make([]StateChanger, len(opsS))

	for i, op := range opsS {
		op, err := parseSingleOperation(op)
		if err != nil {
			return nil, "", "", err
		}

		ops[i] = op
	}

	if len(opsS) == 1 {
		return ops[0], by, signature, nil
	}

	return MultiOp{Ops: ops}, by, signature, nil
}

func parseSingleOperation(input string) (StateChanger, error) {
	parts := strings.Split(input, " ")

	switch parts[0] {
	case "➕":
		name, err := fromAnchor(parts[2])
		if err != nil {
			return nil, errNotAnchor
		}

		return AddParticipant{Name: name, PublicKey: parts[6]}, nil
	case "🪓":
		name, err := fromAnchor(parts[2])
		if err != nil {
			return nil, errNotAnchor
		}

		split, err := strconv.Atoi(parts[4])
		if err != nil {
			return nil, fmt.Errorf("parsing new split value <%s>: %w", parts[4], err)
		}

		return SplitParticipant{Name: name, NewSplit: split}, nil
	case "💻":
		return Configure{NewConfig: parts[3]}, nil
	case "💸":
		name, err := fromAnchor(parts[1])
		if err != nil {
			return nil, errNotAnchor
		}

		amount, err := strconv.Atoi(parts[3][1:])
		if err != nil {
			return nil, fmt.Errorf("parsing $ amount value <%s>: %w", parts[4], err)
		}

		return Spend{Who: name, Amount: amount}, nil
	case "📩":
		from, err := fromAnchor(parts[1])
		if err != nil {
			return nil, errNotAnchor
		}

		to, err := fromAnchor(parts[5])
		if err != nil {
			return nil, errNotAnchor
		}

		amount, err := strconv.Atoi(parts[3][1:])
		if err != nil {
			return nil, fmt.Errorf("parsing $ amount value <%s>: %w", parts[3], err)
		}

		return Transfer{From: from, To: to, Amount: amount}, nil
	}

	return nil, errNotRecognized
}
