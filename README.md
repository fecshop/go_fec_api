<p>
  <a href="http://fecshop.appfront.fancyecommerce.com/">
    <img src="http://img.appfront.fancyecommerce.com/custom/logo.png">
  </a>
</p>
<br/>


Fecshop Go Base Data Api


> Fecshop的底层数据获取部分，通过api，从go实现的api中获取数据，也即是该
> 库包要做的事情，go在数据库连接方面，天生支持数据库连接池，在做数据库连接等底层
> 更加的优秀，这样应对企业级应用高并发，以及为将来的分库分表
> 等方式，做好数据库底层等。

安装
--------

```
go get github.com/fecshop/go_fec_api 
```


配置
----------

进入文件夹：`github.com/fecshop/go_fec_api/config/` 

新建文件 `config.go` ，将 `config_example.go` 的内容复制进去，
将1,2行 和最后一行的注释部分去掉

然后，设置里面相应的参数即可



