#!/bin/sh

candig_repo add-referenceset candig-example-data/registry.db /moredata/human_g1k_v37.fasta
for file in /moredata/*.vcf.gz
do
    candig_repo add-variantset -R human_g1k_v37 candig-example-data/registry.db mock_data HG HG "$file"
done
