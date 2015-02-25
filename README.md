cleo
===========================

## Golang implementation inspired by LinkedIn's Cleo search

### Originally from [https://github.com/jamra/gocleo/](https://github.com/jamra/gocleo/). Areas improved:

* Documents can now be added directly to the index (previously it was limited to being loaded from a text file).
* Document identifiers are now strings instead of integers.
* Project is now more structured/reusable. Previously the package was held together by induvidual functions, now they are organized around a single structure (cleo.Index).

The Cleo search is explained here: [LinkedIn original article](http://engineering.linkedin.com/open-source/cleo-open-source-technology-behind-linkedins-typeahead-search)

### Documentation

[https://godoc.org/github.com/typerandom/cleo](https://godoc.org/github.com/typerandom/cleo)

### Instructions

Sample app:

    package main
    
   	import "github.com/typerandom/cleo"
  
   	func main(){
   	  index := cleo.NewIndex()
   	  
   	  index.LoadFromFile("w1_fixed.txt")
   	  
   	  cleo.ListenAndServe(index, 8080)
    }

Run the program and navigate to http://localhost:8080/search?query=%query%

Where %query% is your search. E.g. try "tractor", "nightingale" or "pizza".

Ex. response for query "pizza":

    [{
        "id": "329017",
        "value": "pizza",
        "score": 1
    }, {
        "id": "329034",
        "value": "pizzaz",
        "score": 0.7142857142857143
    }, {
        "id": "329033",
        "value": "pizzas",
        "score": 0.7142857142857143
    },
    (truncated example...)

### Setup

This should work with go get

    go get github.com/typerandom/cleo
    
### TODO

 - Ability to remove words from index.
 - Add unit tests.
