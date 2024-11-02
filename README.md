# DISAG xml

This library is intended for parsing the result.xml files produced by the OpticScore server by DISAG. Since the
documentation of the format is sparse, some elements have been reverse engineered.

This project is not affiliated with or endorsed by DISAG.

## Structure

The naming of the fields is taken from the XML structure and might not match colloquial use of the terms.

``` 
Result      contains meta information about the file and source system
|
+- Shooters collection of participants participating in a specific discipline (e.g. weapon and age bracket)
   |
   +- Shooter   contains infromation about the partiipant (name, club, etc) 
      |
      +- Shots  collection of Series items
         |
         +- Series  collection of Shots, usually all shot on the same date
            |
            +- Shot a single shot with information about its scoring
```

## Example usage

```go
package main

import (
	"encoding/xml"
	"github.com/dbx12/disag-xml/result"
	"os"
)

func readResultXml(filePath string) (*result.Result, error) {
	bts, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	r := &result.Result{}

	err = xml.Unmarshal(bts, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
```