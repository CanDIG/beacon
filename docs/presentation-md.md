---
layout: presentation
---

# Beacon API

Analysis and shortcomings to adoption

---

## Initial Plan

- simple wrapper around candig server
- just map existing beacon queries to our api calls

## Challenges to Initial Plan

- /search and /count exposing internals that are undocumented
- lack of variant data seeded into development database
- federation wrapper
- inconsistently present next_page_token
- openapi spec gives keys different naming convention
- stored data is read into memory from database
- missing assembly id
