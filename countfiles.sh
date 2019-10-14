#!/bin/bash

echo $(find . -type f | grep -v .git | wc -l)
