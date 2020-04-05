# Caesar Cipher Solver
A Caesar cipher solver that makes calls to an API to check the validity of words after each shift attempt.

My Caesar cipher solver will be a web app. Additionally, to fulfil the objectives of this course and to get an understanding of applied cyber security concepts, I will add DoS protection to the webapp. This will be demonstrated by running a python http server and writting a Go script to query the website many times a second. I will then go on the website from a different device and show that only the IP where the spam is coming from, is blocked.

## Milestones
1. Single word Caesar cipher solver. The word will have to be from the dictionary as an API will be used to match for known words.
2. Multi word Caesar cipher solver. If not all words can be found in the dictionary, return up to two possible matches.
3. Demonstrate a DoS attack on the web app and research mitigation methods, if feasible, implement them.
4. Research and, if feasible, implement a more advanced encryption/decryption techinque.

## Criteria
| Grade | Criteria |
| --- | --- | 
| PS | Created a usable multi-word Caesar cipher solver that only works with words found in an English dictionary. |
| CR | Created a user-friendly multi-word Caesar cipher solver that works in most cases. At least one word has to be found in an English dictionary. |
| DN | Replicate a DoS attack on a website containing a multi word Caesar cipher solver. Research and/or implement DoS mitigation methods. Complete Milestones 1, 2 and 3. |
| HD | Research more advanced encryption/decryption techniques (e.g. RSA, DES, AES) and attempt to implement one. |