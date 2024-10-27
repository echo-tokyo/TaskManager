cd /app/backend

python3 manage.py migrate --noinput

gunicorn --workers=4 --bind 0.0.0.0:8000 back.wsgi
<<<<<<< HEAD
=======

>>>>>>> django
