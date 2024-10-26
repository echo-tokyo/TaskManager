from django.shortcuts import render

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

from django.core.cache import cache
from django.core.mail import send_mail
from django.contrib.auth.tokens import default_token_generator

from users.models import CustomUser
from users.serializers import CustomUserSerializer
from rest_framework_simplejwt.tokens import RefreshToken

# Create your views here.


class CustomUserAPIView(APIView):

    def get(self, request):
        users = CustomUser.objects.all()
        serializer = CustomUserSerializer(users, many=True)
        return Response(serializer.data)

    @staticmethod
    def post(request):
        serializer = CustomUserSerializer(data=request.data)
        if serializer.is_valid():
            user = serializer.save()

            # Генерируем токен для нового пользователя
            refresh = RefreshToken.for_user(user)

            # Возвращаем токены в ответе
            return Response({
                'refresh': str(refresh),
                'access': str(refresh.access_token),
            }, status=status.HTTP_201_CREATED)

        else:
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)