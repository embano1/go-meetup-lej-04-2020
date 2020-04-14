# Deduping with Singleflight

## about://singleflight.Group

> ```go
> type Group struct {
>   // contains filtered or unexported fields
> }
> ```
>
> Group represents a class of work and forms a namespace in which units of work
> can be executed with duplicate suppression.