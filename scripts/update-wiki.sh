#!/bin/bash

cwd="$(pwd)"
pushd "/tmp"
rm -Rf 'terraform-provider-cf.wiki'
git clone https://github.com/cloudfoundry-community/terraform-provider-cf.wiki.git
pushd terraform-provider-cf.wiki
cat <<EOF >_Sidebar.md
## Index ##

* [Provider](provider_config)

### Data Sources ###

EOF
for file in $cwd/website/docs/d/*.markdown; do
  finalname=$(basename $file | sed 's/\.html\.markdown//')
  sed '/^---/,/---/d' "$file" >"datasource_${finalname}.md"
  echo "* [cloudfoundry_${finalname}](datasource_${finalname})" >>_Sidebar.md
done
echo "" >>_Sidebar.md
echo "### Resources ###" >>_Sidebar.md
echo "" >>_Sidebar.md

for file in $cwd/website/docs/r/*.markdown; do
  finalname=$(basename $file | sed 's/\.html\.markdown//')
  sed '/^---/,/---/d' "$file" >"resource_${finalname}.md"
  echo "* [cloudfoundry_${finalname}](resource_${finalname})" >>_Sidebar.md
done
sed '/^---/,/---/d' $cwd/website/docs/index.html.markdown >provider_config.md

git add .
git commit -m "update wiki"
git push origin master
popd
rm -Rf 'terraform-provider-cf.wiki'
popd
