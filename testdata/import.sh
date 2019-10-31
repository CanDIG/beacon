#!/bin/sh

candig_repo list /data/candig-example-data/registry.db
candig_repo add-referenceset /data/candig-example-data/registry.db /moredata/human_g1k_v37.fasta
for file in /moredata/*.vcf.gz
do
    candig_repo add-variantset -R human_g1k_v37 /data/candig-example-data/registry.db mock_data HG HG "$file"
done
