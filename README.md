# alchemy
An acceptance testing tool for Algolia indexes

## Intro

Algolia's index settings are flexible, interrelated, and very easy to change. It can be tempting to fix small relevance problems by tweaking custom ranking or query rules settings, and break results elsewhere.

This is where `alchemy` comes in: you can now run acceptance tests against a set of index settings.



## Go Environment setup

Get Go and [dep](https://github.com/golang/dep) installed:

    $ brew install go dep

Run `dep ensure` to check out the dependencies:

    $ dep ensure

Install and run:

    $ go install
    $ alchemy index_name

Or just run:

    $ go run . index_name


## How to Use

Create a blank index called `test_index`.

You define a set of sample records (in JSON format). Create a file called `fixtures.json`:

    [
      {
        "objectID": "beuller",
        "name": "Ferris Beuller's Day Off",
        "leads": ["Matthew Broderick", "Alan Ruck", "Mia Sara"],
        "year": 1986,
        "box_office": 70136369
      },
      {
        "objectID": "b2tf",
        "name": "Back to the Future",
        "leads": ["Michael J. Fox", "Christopher Lloyd", "Lea Thompson"],
        "year": 1985,
        "box_office": 210609762
      },
      {
        "objectID": "pulp_fiction",
        "name": "Pulp Fiction",
        "leads": ["John Travolta", "Uma Thurman"],
        "year": 1994,
        "box_office": 107928762
      },
      {
        "objectID": "faceoff",
        "name": "Face/Off",
        "leads": ["John Travolta", "Nicholas Cage"],
        "year": 1997,
        "box_office": 112225777
      }
    ]

It's better to use something human-friendly for your fixtures' `objectID`; we'll use them to specify expected results in a moment.

We also need to specify the tests and expected results. Create a file called `index_name.test.json`:

    [
      {
        "query": "off",
        "expectedResults": [ "faceoff", "beuller" ]
      },
      {
        "query": {
          "filters": "year < 1990"
        },
        "expectedResults": [ "b2tf", "beuller" ]
      },
      {
        "query": {
          "query": "off",
          "filters": "year < 1990"
        },
        "expectedResults": [ "beuller" ]
      }
    ]

Configure the `alchemy` tool with an `.alchemyrc` file in your current directory (or specify it with the `-c` cli flag):

    {
      "appId": "algolia app ID here",
      "searchKey": "algolia search key here",
      "secretKey": "algolia secret key here",
      "fixtures": "./fixtures.json",
      "tests": "./index_name.test.json"
    }

...and run the tool against the index to see the results:

    $ alchemy index_name
        ✔ "off"
        ✖ "" [year < 1990]
                - Expected results weren't returned
        ✔ "off" [year < 1990]

Try adding `box_office` as a descending custom ranking attribute and run it again:

    $ alchemy index_name
        ✔ "off"
        ✔ "" [year < 1990]
        ✔ "off" [year < 1990]


## To Do

- [ ] Rewrite in JavaScript
- [x] Coloured output
- [ ] Better error output
- [ ] Filters in query object
- [ ] _Tests_
- [ ] Advance query rule tests
- [ ] More validation for `indexName` (ensure its length is < 256 - timestamp length (10) - `len('alchemy__')`)
