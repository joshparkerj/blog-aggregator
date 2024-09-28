#! /bin/bash
cd sql/schema || exit
goose postgres postgres://postgres:postgres@localhost:5433/gator down-to 0
goose postgres postgres://postgres:postgres@localhost:5433/gator up
