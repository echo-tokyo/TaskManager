# Generated by Django 5.0.6 on 2024-10-26 09:20

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('users', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='customuser',
            name='job_title',
            field=models.CharField(max_length=50),
        ),
    ]
