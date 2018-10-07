# wordspacesaver
The premise of this concept is that words are repeated and take up space. Make a list of common words, repalce them with a utf character the represents the word's place on that map, and it will save space.

Here is how it works:

  Read list of most common words and store their order value in an array, and the word value with the order value in a map.
  
  To encode:
    Read input string and if word is among the stored values in the map, replace it with the utf character whose decimal value represents the word. ex map[word]int
      There is also one letter before it, thus the word now only takes up 2-4 bytes. ex @€
  
  To decode:
    Read encoded string, check for @, check @ word size, convert utf symbol to int value, find value in the array, replace symbol with stored word.
 
Here is a working example: https://play.golang.org/p/OfkSZXhgfpv
