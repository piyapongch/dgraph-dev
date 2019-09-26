# gdb2dg
This command converts GraphDB export in n-quads to Dgraph live import data file.
- Add ```<dgraph.type>``` "IRObject" to all subjects
- Add ```<dgraph.type>``` for ```<info:fedora/fedora-system:def/model#hasModel>``` predicates
- Add ```<xid>``` predicates for all subjects, Dgraph will replace the subjects with generated uids.
- Convert IRI object to string for these predicates:
  - ```<http://www.loc.gov/premis/rdf/v1#hasMessageDigest>```
  - ```<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>```
  - ```<http://www.w3.org/ns/ldp#hasMemberRelation>```
  - ```<http://www.w3.org/ns/ldp#insertedContentRelation>```
  - ```<http://fedora.info/definitions/v4/repository#hasFixityService>```
  - ```<http://www.iana.org/assignments/relation/describedby>```