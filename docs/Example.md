# Example SP

## At a glance [(ℹ)]((https://github.com/jazcarate/sp/blob/master/docs/understanding_a_report.md.md#at-a-glance))
| | [Joaco](#joaco) | [Greenderella](#greenderella) | [Bob](#bob) |
| --- | ---: | ---: | ---: |
| [Joaco](#joaco) | | $5 ↗ | $7 ↙ |
| [Greenderella](#greenderella) | $5 ↙ |  | $10 ↙ |
| [Bob](#bob)  | $7 ↗ | $10 ↗ |  |


## Participants
| Name | Split | %age |
| --- | ---: | ---: |
| [Joaco](#joaco) | 1 | 50% |
| [Greenderella](#greenderella) | 1 | 50% |
| [Bob](#bob)  | 0 |


## Details
### Joaco
<!-- Public key 1nrNwqN/3ssdTwJU4b8LKml1dBwhf2lV49AF2Vqf1Y9Uhnwy5+zF0DfDC0B8ZvSy3eDN+wLu5IUYHZY+XLBbvZJvUWl6VLznp6g3P3thlnF0LAe//lHrY0opuLSa+vIfi8gTBTtG9w1cdPyzhzjUCVHMqAH82I9+NywekJJI3w3IbO3pdvGm5uuKBoapj2LxDgzNqnPdQH8dTmeKmxfXsoAJLF1x9+UUSEhsZv03hWBwDq8jW1NlhzUWYeGKDP7Z3YtV1j0cTWyVZKws4FOfohahVDE9pPPI4fMkW6NJD2JzL4ZeCmgrSdxqWCLGIpZq95oMTtMXqkKuTetG+KkmqQ== -->
You currently split the expenditures at **50%**.
- You **owe $5** to [Greenderella](#greenderella)
- You are **owed $7** from [Bob](#bob)

### Greenderella
<!-- Public key 6P8mIpUnW+9kN4Gm9tccCHfTo9VGR8jxOTMBpk0qHwrLdwXR6aYVJssHS00yrPTk1h9sJpCGuMiGQTHhrIxT4UasFvw2GW4CLzS+n/vmyzAsrn58jzub8VN/DJ/PhmfdNK5Y+ty2veSRD+ShMhP3k9DL5gc9tAHgextiUaFp8retVCFepZKCJR0b3iuO98+RBvVbAh765qLzX5tP2Klcm8vAKdq1RZe3butVSrxGo9j0UqJGpoA4xdoEpsXIxnKl9hy4cswDU5qZo+z7GJjz1G78EXMxCDzU7iy0+xXNg5e3n2wfhpbb0prBR13lvqKYhwg5ESBPC8R9bH8rJ7v9Lw== -->
You currently split the expenditures at **50%**.
- You are **owed $5** from [Joaco](#joaco)
- You are **owed $10** from [Bob](#bob)

### Bob
<!-- Public key 6no8Rv4VaEn9NEDmETm49qOdFawwVZU+i9H/IGkuGv6Gn3+0bYMsKyxo00eCOvNikmiXUojvR0/O2Gh6/OXc4H1+/CAiNEpNM67PsRtlYZ5gQUxIWsD0bA64Jr/bJcZX2gHKZXtGT0ezkpLUVxZ1UpZkvngDyfV4ouF4t9u4J0+nkXREEqGKDkxd9cKj2u9aCnoEveYd2iKvYtAFiuBcyLFeJJaORmznrh0Qh9SgGeM8vrd39GSkZIqOw1YPQN93uWdB6CuIucFKDjrRsv8CGFZ5GRTPyyqR/PmO2l75ZQwdyahIjtFWZzDy6lFZispclxX84uEXqVhlGrVmcBSb/Q== -->
You currently split the expenditures at **0%**.
- You **owe $10** to [Greenderella](#greenderella)
- You **owe $7** to [Joaco](#joaco)

## Operations
Current trust configuration: **Trust** [(ℹ)](https://github.com/jazcarate/sp/blob/master/docs/understanding_a_report.md.md#trust)

### Log
| # |  Operation | On | Note | Status |
| ---: | --- | --- | --- | ---: |
| [5](#op-5)<a id="op-5"></a> | 📩 [Joaco](#joaco) **transferred $7** to [Bob](#bob)<!-- Sign &joaco 1-->  |2021-03-2T18:00:15+01:00 | | ✅ |
| [4](#op-4)<a id="op-4"></a> | 🪓 Split [Bob](#bob) for `0`<!-- Sign &joaco 1-->  | 2021-03-2T17:50:20+01:00 | | ✅ |
| [3](#op-3)<a id="op-3"></a> | 📩 [Joaco](#joaco) **transferred $5** to [Greenderella](#greenderella)<!-- Sign &joaco 1--> | 2021-03-2T17:47:20+01:00 |  _Sorry for the delay_ | ✅ |
|[2](#op-2)<a id="op-2"></a> | 💸 [Greenderella](#greenderella) **spent $30**<!-- Sign &greenderella 1--> |2021-03-2T17:46:10+01:00 |  _Amazon 📦📦_  | ✅ |
|[1](#op-1)<a id="op-1"></a> | ➕ Add [Joaco](#joaco)<!-- Public key: 1nrNwqN/3ssdTwJU4b8LKml1dBwhf2lV49AF2Vqf1Y9Uhnwy5+zF0DfDC0B8ZvSy3eDN+wLu5IUYHZY+XLBbvZJvUWl6VLznp6g3P3thlnF0LAe//lHrY0opuLSa+vIfi8gTBTtG9w1cdPyzhzjUCVHMqAH82I9+NywekJJI3w3IbO3pdvGm5uuKBoapj2LxDgzNqnPdQH8dTmeKmxfXsoAJLF1x9+UUSEhsZv03hWBwDq8jW1NlhzUWYeGKDP7Z3YtV1j0cTWyVZKws4FOfohahVDE9pPPI4fMkW6NJD2JzL4ZeCmgrSdxqWCLGIpZq95oMTtMXqkKuTetG+KkmqQ== --><!-- Sign &joaco 1--><br>🪓 Split [Joaco](#joaco) for `1`<!-- Sign &joaco 1--><br>➕ Add [Greenderella](#greenderella)<!-- Public key: 6P8mIpUnW+9kN4Gm9tccCHfTo9VGR8jxOTMBpk0qHwrLdwXR6aYVJssHS00yrPTk1h9sJpCGuMiGQTHhrIxT4UasFvw2GW4CLzS+n/vmyzAsrn58jzub8VN/DJ/PhmfdNK5Y+ty2veSRD+ShMhP3k9DL5gc9tAHgextiUaFp8retVCFepZKCJR0b3iuO98+RBvVbAh765qLzX5tP2Klcm8vAKdq1RZe3butVSrxGo9j0UqJGpoA4xdoEpsXIxnKl9hy4cswDU5qZo+z7GJjz1G78EXMxCDzU7iy0+xXNg5e3n2wfhpbb0prBR13lvqKYhwg5ESBPC8R9bH8rJ7v9Lw== --><!-- Sign &joaco 1--><br>🪓 Split [Greenderella](#greenderella) for `1`<!-- Sign &joaco 1--><br>➕ Add [Bob](#bob)<!-- Public key: 6no8Rv4VaEn9NEDmETm49qOdFawwVZU+i9H/IGkuGv6Gn3+0bYMsKyxo00eCOvNikmiXUojvR0/O2Gh6/OXc4H1+/CAiNEpNM67PsRtlYZ5gQUxIWsD0bA64Jr/bJcZX2gHKZXtGT0ezkpLUVxZ1UpZkvngDyfV4ouF4t9u4J0+nkXREEqGKDkxd9cKj2u9aCnoEveYd2iKvYtAFiuBcyLFeJJaORmznrh0Qh9SgGeM8vrd39GSkZIqOw1YPQN93uWdB6CuIucFKDjrRsv8CGFZ5GRTPyyqR/PmO2l75ZQwdyahIjtFWZzDy6lFZispclxX84uEXqVhlGrVmcBSb/Q== --><!-- Sign &joaco 1--><br>🪓 Split [Bob](#bob) for `1`<!-- Sign &joaco 1--> |2021-03-2T17:45:56+01:00 | | ✅ |