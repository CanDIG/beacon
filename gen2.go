// +build generate
package main
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/bio_metadata_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/sequence_annotation_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/references.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/rna_quantification_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/genotype_phenotype.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/metadata_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/sequence_annotations.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/pipeline_metadata_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/read_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/pipeline_metadata.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/allele_annotations.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/genotype_phenotype_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/variant_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/clinical_metadata_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/rna_quantification.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/search_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/variants.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/genotype_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/reads.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/peer_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/clinical_metadata.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/bio_metadata.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/common.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/reference_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/allele_annotation_service.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/candig/metadata.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/google/api/annotations.proto
//go:generate protoc -I candig-schemas/src/main/proto --go_out=./go/ ./candig-schemas/src/main/proto/google/api/http.proto
