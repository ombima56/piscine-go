#!/bin/bash
curl https://learn.zone01kisumu.ke/assets/superhero/all.json |jq "[map(select(.id == ${HERO_ID}))] | .[] | .[] .connections.relatives" | tr -d '"'