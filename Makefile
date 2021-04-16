pull:
	scp hgs:apps/coved/data.json .

push:
	echo "are u sure?"
	read
	scp data.json hgs:apps/coved/data.json
