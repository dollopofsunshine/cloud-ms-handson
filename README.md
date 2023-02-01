# Cloud and Microservices - Hands-on

1. We have a golang container that serves a webpage. It is built using the attached dockerfile. Security audit has revealed that the container contains source code for the application which is insecure. How will you modify the dockerfile to ensure that the final container contains only the binaries? Suggest 2 ways

     *Hint: Go is a compiled language.*

2. Create a container using the attached Dockerfile. Build the image and note down the image size. Study the commands in the file and identify what can be done to reduce the image size. What is the image size of the optimised container?

3. You have a containerized application that renames the file present in the user's home directory. Modify the below docker file to allow the code to work using the ~/demo-files directory as the source.
 
     *Hint: Create a volume using docker create. Run a container and mount the volume.*

 4. In the previous case, how will you modify the same container such that it will copy the source files to ~/demo-result. You should ensure that your container has only read permissions on ~/demo-files as it is a security requirement.

     *Hint: Mount the volume as read-only.*

 5. This container runs as the root user. Reviewer find that your set-up is not following the principle of least privileges. How will you modify the container to run it with a lower privileged user named "customer"?

     *Hint: Add a new user to the container using adduser in Dockerfile and set the user as default user*

 6. You find that a new image is added to your container registry. You do not have the Dockerfile used to build the image. How will you verify the commands used to build this image without executing it? How will you find the start-up command for this container?

     *Command to pull the Image: docker pull dollopofsunshine/handsonexerciseamd64 --platform=linux/amd64*

