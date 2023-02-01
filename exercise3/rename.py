import os
path = $HOME
for filename in os.listdir(path):
  os.rename(path+'/'+filename+'.txt', path+'/'+filename+'.bak')
