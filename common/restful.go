// Copyright (c) 2012-2013 Aslan Varoqua.  This is Free Software, released
// under the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for
// details.  Resist intellectual serfdom - the ownership of ideas is akin to
// slavery.

// I use  package napping, with HTTP Basic
// authentictation over HTTPS, to retrieve a Github auth CosmicJs lambda.
package gocosmic

import "errors"

import (
	"fmt"
	"gopkg.in/jmcvetta/napping.v3"
	"log"
	"net/url"
	"time"
  "github.com/aslanvaroqua/gocosmic/common"
  )
  

type Common struct { }

func Common(url) Get {
      //
      // Send request to server
      //
      resp, err := s.Get(url,&res, &e)
      if err != nil {
        log.Fatal(err)
      }
}

func Common(url) Get {
      //
      // Send request to server
      //
      resp, err := s.Get(url,&res, &e)
      if err != nil {
        log.Fatal(err)
      }
}

func Common(url,payload) Post {
      //
      // Send request to server
      //
      resp, err := s.Post(url,&payload, &res, &e)
      if err != nil {
        log.Fatal(err)
      }
}
