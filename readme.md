rmcrashers is a tool to help with go-fuzz triage.

go-fuzz can find a whole bunch of duplicate crashers in a short time.

rmcrashers lets you provide a regexp and then deletes all entries
from the crashers directory for which the output matches the regexp.
