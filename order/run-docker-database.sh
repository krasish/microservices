docker run -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=verysecretpass \
  -e MYSQL_DATABASE=order mysql