#!/bin/sh

cd "$(dirname "$0")" || exit

for ref in ./*.fasta.gz
do
    gunzip -k "$ref"
done
for varset in ./*.vcf.gz
do
    tabix -p vcf "$varset"
done
