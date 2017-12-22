package doctype

type Engine interface{
    Type () string
    String () string
}
