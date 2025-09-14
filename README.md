handler -> build, interact with external service (FE, client CLI, kafka, ...) - request grpc - request http - tieeps nhanaj event tu kafka - tiep nhan trigger manually tu cronjobs (CLI)

- can khoi tao, goi ham New o dau ->trong qua trinh khoi tao project -> cam vao server grpc cua minh -> khoi taoj new ben ngoai logic -> dependency injection -> wire -> tu dong nha biet cam dataaccessor vao dau

Task to do left:

1. Add protocvalidate - got some problems with buf cli to update
2. Cleanup code - the function name and params not really consistent through project and
