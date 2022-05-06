# Reckoner
A 🚀***BLAZING FAST***🚀CLI data parser and explorer. Parse, query, or explore data dynamically.

## How to use

Without flags: Prints all values found with given keys  
Search flags: For conditional searches, use -s=searchTerm and -sl=keyIfFound with keys as args.  
Examples:  
To find all values found with "title" key  
`pipe | reckoner title`<br><br>
To find "url" if key with "title" whose value is "hello" is found  
`pipe | reckoner -s=hello -sl=url title`  
