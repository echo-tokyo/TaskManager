from django.db import models
from users.models import CustomUser


# Create your models here.


class Board(models.Model):
    title = models.CharField(max_length=50)


class Task(models.Model):
    board = models.ForeignKey(Board, on_delete=models.CASCADE)
    title = models.CharField(max_length=50)
    task_desc = models.TextField()
    status = models.CharField(max_length=20)
    created_at = models.DateTimeField(auto_now_add=True)
    executors = models.ManyToManyField(CustomUser)


class Comment(models.Model):
    user = models.ForeignKey(CustomUser, on_delete=models.CASCADE)
    task = models.ForeignKey(Task, on_delete=models.CASCADE)
    comment_desc = models.TextField()
