from django.http import HttpResponse
from django.shortcuts import render
from rest_framework.response import Response
from rest_framework.views import APIView

# Create your views here.


class PingView(APIView):

    def get(self, request):
        print(request.user)
        if request.user.is_active:
            return HttpResponse("Pong")
        else:
            return HttpResponse("Not active")