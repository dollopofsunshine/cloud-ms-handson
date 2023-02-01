FROM ubuntu:18.04 

WORKDIR /root
COPY . .

CMD [ "ls", "source"]
CMD [ 'for file in `ls source/*.txt`; do; mv "$file" "$file.bak"; done;']
CMD [ "ls", "source"]