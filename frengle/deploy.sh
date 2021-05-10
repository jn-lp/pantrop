#!/bin/bash

set -e

git pull origin development
git checkout production
git merge development
git push origin production
git checkout development
