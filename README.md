# pattern matching evaluations with functional encryption

This project develops an example of pattern matching with functional encryption. A finite automata is employed as algorithm for pattern searching.
A client encrypts pattern and text and sends them out to the decryptor. Then it ouputs the matches and send them back to the client.

The adopted schema is [FHIPE](https://eprint.iacr.org/2016/440.pdf) (Function Hiding Inner Product Encryption), since it permits of encrpypting both text and pattern. 
The functional encryption library is [Gofe](https://github.com/fentec-project/gofe).


## Diagram

![](./pattern_matching.png) 

## Files
- [PatternMatching.py](./patternMatching.py): it is a workaround file for creating matrices of text and pattern that are used by the functional encryption application. 
- [PatternMatching.go](./PatternMatching.go): it contains the FHIPE schema, retrieves pattern and text matrices for the encryption and decryption task.   

## ISSUES
* With the current inner product encrption schema the encryption and decryption time is too slow for long texts. Possible Solutions are: 
  * Check if other solutions exist in letterature in searchable encryption field. 
  * Optimize/replace the current pattern searching algorithm, like with a suffix tree. 
  * Try with another inner product schema.  


