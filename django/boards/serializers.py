from rest_framework import serializers
from .models import *
from users.serializers import CustomUserSerializer


class BoardSerializer(serializers.Serializer):
    id = serializers.IntegerField(read_only=True)
    title = serializers.CharField()

    def create(self, validated_data):
        board = Board.objects.create(**validated_data)
        board.save()
        return board


class TaskSerializer(serializers.Serializer):
    id = serializers.IntegerField(read_only=True)
    board_id = serializers.IntegerField(allow_null=True, required=True) # Представляем как ForeignKey
    title = serializers.CharField(max_length=50)
    task_desc = serializers.CharField(allow_blank=True)
    status = serializers.CharField(max_length=20)
    created_at = serializers.DateTimeField(read_only=True)
    executors = CustomUserSerializer(many=True)

    def create(self, validated_data):
        executors_data = validated_data.pop('executors')
        task = Task.objects.create(**validated_data)
        for executor_data in executors_data:
            executor, created = CustomUser.objects.get_or_create(**executor_data)
            task.executors.add(executor)
        return task

    def update(self, instance, validated_data):
        executors_data = validated_data.pop('executors', None)
        instance.title = validated_data.get('title', instance.title)
        instance.task_desc = validated_data.get('task_desc', instance.task_desc)
        instance.status = validated_data.get('status', instance.status)
        instance.save()

        if executors_data is not None:
            instance.executors.clear() # Очистка текущих исполнителей
            for executor_data in executors_data:
                executor, created = CustomUser.objects.get_or_create(**executor_data)
                instance.executors.add(executor)

        return instance

    def to_representation(self, instance):
        # Получаем стандартное представление объекта
        representation = super().to_representation(instance)
        # Извлекаем только дату из created_at
        representation['created_at'] = instance.created_at.date()
        return representation
