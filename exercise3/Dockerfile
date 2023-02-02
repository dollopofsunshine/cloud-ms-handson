FROM ubuntu:18.04 

WORKDIR /root
COPY . .

#CMD [ 'for file in `ls source`; do; mv source/"$file" source/"$file.bak"; done;']
CMD ["/bin/bash", "-c", "rename.sh"] 
CMD [ "ls", "source"]