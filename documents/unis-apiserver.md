# unis-apiserver

Image Registry:  

Public images: 
apiserverIP:PORT:/images/public/images  
apiserverIP:PORT:/images/public/{imageID}  
Private images: 
apiserverIP:PORT:/images/{username}/images  
apiserverIP:PORT:/images/{username}/{imageID}  

Account management:  

apiserverIP:PORT:/users

Node management:

Public nodes: 
apiserverIP:PORT:/nodes/public/nodes  
apiserverIP:PORT:/nodes/public/{nodeID}  
Private images: 
apiserverIP:PORT:/nodes/{username}/nodes  
apiserverIP:PORT:/nodes/{username}/{nodeID}  

Instance Management:
apiserverIP:PORT:/instances/{username}/instances  
apiserverIP:PORT:/instances/{username}/{instanceID}  