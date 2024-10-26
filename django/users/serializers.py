from rest_framework import serializers
from users.models import CustomUser


class CustomUserSerializer(serializers.Serializer):
    id = serializers.IntegerField(read_only=True)
    username = serializers.CharField(max_length=20)
    password = serializers.CharField(write_only=True)
    fio = serializers.CharField(max_length=50)
    is_staff = serializers.BooleanField(default=False)
    job_title = serializers.CharField(max_length=50)

    def create(self, validated_data):
        user = CustomUser.objects.create_user(**validated_data)
        return user

    def update(self, instance, validated_data):
        for key, value in validated_data.items():
            if hasattr(instance, key):
                setattr(instance, key, value)

        instance.save()
        return instance