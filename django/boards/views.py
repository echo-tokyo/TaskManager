from django.http import HttpResponse
from django.shortcuts import render
from rest_framework import status
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response
from rest_framework.views import APIView

from boards.models import Task, Board
from boards.serializers import TaskSerializer, BoardSerializer


# Create your views here.


class PingView(APIView):

    def get(self, request):
        print(request.user)
        if request.user.is_active:
            return HttpResponse("Pong")
        else:
            return HttpResponse("Not active")


class BoardView(APIView):
    permission_classes = (IsAuthenticated,)

    def get(self, request):

        board = Board.objects.all()

        serializer = BoardSerializer(board, many=True)

        return Response(serializer.data)

    def post(self, request):

        serializer = BoardSerializer(data=request.data)

        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        else:
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class TaskAPIView(APIView):
    permission_classes = (IsAuthenticated, )

    def get(self, request, *args, **kwargs):

        pk = kwargs.get('pk', None)
        if pk:
            task = Task.objects.get(pk=pk)
            serializer = TaskSerializer(task)
            return Response(serializer.data)
        tasks = Task.objects.all()

        backlog = []
        proccessing = []
        finished = []

        serializer = TaskSerializer(tasks, many=True)

        for task in serializer.data:
            if task['status'] == 'finished':
                proccessing.append(task)
            elif task['status'] == 'proccessing':
                finished.append(task)
            else:
                backlog.append(task)

        response_data = {
            'backlog': backlog,
            'proccessing': proccessing,
            'finished': finished,
        }

        # Возвращаем ответ
        return Response(response_data, status=status.HTTP_200_OK)

    def post(self, request):
        serializer = TaskSerializer(data=request.data)

        if serializer.is_valid():

            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        else:
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    def patch(self, request, pk):
        task = Task.objects.get(pk=pk)
        serializer = TaskSerializer(task, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

    def delete(self, request, pk):
        task = Task.objects.get(pk=pk)

        task.delete()

        return Response(status=status.HTTP_200_OK)
