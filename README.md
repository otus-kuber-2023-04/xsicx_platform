# xsicx_platform
xsicx Platform repository

### HOMEWORK 1
1. **Почему все pod в namespace kube-system восстановились
после удаления?** 

- Для **coredns** существует ReplicaSet, что гарантирует  гарантирует, что определенное количество экземпляров подов будет  запушено в мастере `kubectl get replicaset -n kube-system`
- **kube-proxy** - Deamon Set, объект API Kubernetes, который гарантирует, что определенный под будет запущен на всех (или некоторых) узлах `kubectl get daemonset -n kube-system`
- Остальные поды это static pods (https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/)
2. **Выполнение основного задания**
- Установлен kubectl, minikube
- Запущен minikube и проверена устойчивость к отказам
- Создан Dockerfile отвечающий поставленной задаче
- Создан манифест web-pod.yaml и проверена работоспособность
- Рассмотрены основные команды по работе с подом
- Добавлен init контейнер и volumes
- Проверена работоспособность приложения
- Установлен kube-forwarder
3. **Причина, по которой pod frontend находится в статусе Error**
- На ARM64 падала ошибка `qemu-x86_64: Could not open '/lib/ld-musl-x86_64.so.1': No such file or directory`. Проблема пофиксилась изменением докер образа с `FROM golang:1.20.4-alpine@sha256:4ee203ff3933e7a6f18d3574fd6661a73b58c60f028d2927274400f4774aaa41 as builder` на `FROM golang:1.20.4-alpine as builder` в DockerFile
- golang приложение не могло найти env переменные и кидало panic. Проблема пофиксилась добавлением требуемых переменных в yaml файл
