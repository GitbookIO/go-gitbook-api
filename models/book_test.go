package models

import (
	"encoding/json"

	"testing"
)

func TestUnmarshaling(t *testing.T) {
	book := Book{}
	err := json.Unmarshal([]byte(bookJSON), &book)

	if err != nil {
		t.Error(err)
	}
}

const bookJSON string = `
{
    "id":"aaronomullan/simple-book",
    "name":"simple-book",
    "title":"simple book",
    "description":"",
    "public":true,
    "price":0,
    "githubId":"",
    "categories":[

    ],
    "cover":{
        "large":"http://localhost:5000/cover/book/aaronomullan/simple-book?build=1416256968809",
        "small":"http://localhost:5000/cover/book/aaronomullan/simple-book?build=1416256968809"
    },
    "urls":{
        "access":"http://localhost:5000/book/aaronomullan/simple-book",
        "homepage":"http://localhost:5000/book/aaronomullan/simple-book",
        "read":"http://localhost:5000/read/book/aaronomullan/simple-book",
        "reviews":"http://localhost:5000/book/aaronomullan/simple-book/reviews",
        "subscribe":"http://localhost:5000/subscribe/book/aaronomullan/simple-book",
        "download":{
            "epub":"http://localhost:5000/download/epub/book/aaronomullan/simple-book",
            "mobi":"http://localhost:5000/download/mobi/book/aaronomullan/simple-book",
            "pdf":"http://localhost:5000/download/pdf/book/aaronomullan/simple-book"
        }
    },
    "author":{
        "username":"aaronomullan",
        "name":"Aaron O'Mullan",
        "urls":{
            "profile":"http://localhost:5000/@aaronomullan"
        },
        "accounts":{
            "twitter":"AaronOMullan"
        }
    },
    "license":{
        "id":"nolicense",
        "layout":"license",
        "permalink":"http://choosealicense.com/licenses/no-license/",
        "category":"No License",
        "class":"license-types",
        "title":"No License",
        "description":"You retain all rights and do not permit distribution, reproduction, or derivative works. You may grant some rights in cases where you publish your source code to a site that requires accepting terms of service. For example, publishing code in a public repository on GitHub requires that you allow others to view and fork your code.",
        "note":"This option may be subject to the Terms Of Use of the site where you publish your source code.",
        "how":"Simply do nothing, though including a copyright notice is recommended.",
        "required":[
            "include-copyright"
        ],
        "permitted":[
            "commercial-use",
            "private-use"
        ],
        "forbidden":[
            "modifications",
            "distribution",
            "sublicense"
        ],
        "url":"http://choosealicense.com/licenses/no-license/",
        "content":"Copyright [year] [fullname]",
        "path":"licenses/no-license.html"
    },
    "language":{
        "code":"en",
        "name":"English",
        "nativeName":"English"
    },
    "reviews":{
        "count":0,
        "rating":0
    },
    "transactions":{
        "count":0,
        "donations":false
    },
    "dates":{
        "created":"2014-10-13T19:21:03.070Z"
    },
    "permissions":{
        "read":true,
        "write":true,
        "manage":true
    }
}
`
