---
layout: presentation
---

# Beacon API

Analysis and shortcomings to adoption

---

## Initial Plan

- simple wrapper around candig server
--

- just map existing beacon queries to our api calls

![Beacon Diagram](./beacon.svg)
---

## Challenges to Initial Plan
--

- /search and /count exposing internals that are undocumented
--

- lack of variant data seeded into development database
--

- federation wrapper
--

- inconsistently present next_page_token
--

- openapi spec gives keys different naming convention
--

- stored data is read into memory from database
--

- missing assembly id

---

## Conclusions

--

While a Beacon implementation on top of CanDIG is certainly possible.
The current state of `candig-server`'s API imposes significant
challenges to clients building on top of CanDIG.

--

Going forward into V2, adoption of basic schema testing and an OpenAPI
first approach to implementing the server layer will ensure the
documentation doesn't fall out of sync with the implementation as much
as it has in this instance.
