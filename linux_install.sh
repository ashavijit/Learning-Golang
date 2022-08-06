wget https://go.dev/dl/go1.19.linux-amd64.tar.gz
cd /Downloads
tar -xzvf go1.19.linux-amd64.tar.gz
sudo mv go /usr/local
#adding env variables in linux by command not manually
vi ~/.bashrc
#press SHIFT+G to move to the bottom 
#press O to add new line to the next line
#press CTRL+C and :wq 
source ~/.bashrc #to save the written file
