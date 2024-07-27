# Challenge


## Output

Return the output to STDOUT with the following format for each unique station, find the minimum, average and maximum temperature recorded
`{<station name>:<min>/<average>/<max>;<station name>:<min>/<average>/<max>}`


## Facts

Problem constrains:
- the temperature value is within [-99.9, 99.9] range.
- the temperature value only has exactly one fractional digit.
- the byte length of station name is within [1,100] range.
- there will be a maximum of 10,000 unique station names.
- rounding of temperature values must be done using the semantics of IEEE 754 rounding-direction "roundTowardPositive”.

