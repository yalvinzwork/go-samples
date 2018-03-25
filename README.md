go build ./cmd/frontend/ && ./frontend -debug -l /var/log/ipreit/ipreit.access.log -e /var/log/ipreit/ipreit.error.log
go build ./cmd/web/ && ./web -debug -l /var/log/ipreit/ipreit.access.log -e /var/log/ipreit/ipreit.error.log
(without flag -l and -e, log will be printed to stdout, for development purpose they aren't necessary)
