Sending an Encrypted Session

1. Build and run GenerateKey folder. This will generate a sharable public key in this folder, and a server to wait for a session key to be decrypted with it's corsponding pprivate key

2. Build and run GenerateSession folder. This will generate the SessionKey with random data that will then be encypted with the public key in this folder. It then will call to the Generate Key and pass that ciphertext to be decoded. 


Both will echo it's SessionKey in a rune array to prevent any charectors from being misformatted when being written, since they are random bytes and could be any number of charectors, including different types of whitespaces.

Comparing these two rune arrays should be identical
