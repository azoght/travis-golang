# anagrammer
Purpose: takes words and lists all possible rearrangements of them, perfect for scrambling/unscrambling words (JUMBO doesn't recommend)

Functions:

contains (takes string slice and string, returns boolean) This function determines if the string slice argument contains the string argument by iteration. If the string slice contains the string, the function returns true. If not, the function returns false.

factorial (takes integer, returns integer) This function determines the factorial of the integer argument with a for loop. If the integer is negative or a 0, the function returns 0. If not, it calculates and returns the integer's factorial as an integer.

stringToCharArray (takes string, returns string slice and error) This function converts a string to a slice of its seperate characters. If the string contains a space, the conversion fails and returns an error. If not, each character becomes a part of the string slice, and the slice is returned with no error.

listAnagrams (takes string) •This function contains the previous three functions listed.

•It lists all anagrams and the original form of the string.

•First, the string is converted to a slice using stringToCharArray().

•The function also returns an error, so if there is an error, the error is printed, and anagrams can't be listed.

•Otherwise, the function continues to the stage where the number of letters is calculated.

•This is critical for words that repeat the same letter more than once.

•A for loop counts the number of each letter in the string and deducts it from the length of the string slice of letters.

•The loop also adds letters that appear more than once in the string to a seperate string slice, while using the contains() function.

•Afterwards, the function creates a new string slice of the anagrams to be printed, which has an initial value of the string argument.

•It continues to a big loop of for loops and conditional statements.

•The first loop switches letters of the slice a random number of times.

•Then, the rearranged slice is converted to a string, then added to the slice of anagrams, if the slice doesn't contain it.

•The next loop uses a boolean value to check if it's a duplicate.

•Another conditional statement checks if the length of the anagrams slice is equal to the factorial of the number of letters, using the factorial() function.

•If the length of the slice is equal to the calculated factorial, the loop is exited, then each element of the anagram slice is printed.

Next Steps:

I plan to add user input to this program, so that anyone can import their own words and see which anagrams are created. I also want to rearrange the anagram slice so that each element is printed in alphabetical order.

What is an anagram exactly?

Dictionary.com defines an anagram as a word, phrase, or sentence formed from another by rearranging its letters. My program, on the other hand, counts non-sensical words as anagrams. Go ahead and use it for secret code too!

Hope you like this project! BTW, it's written in Golang
