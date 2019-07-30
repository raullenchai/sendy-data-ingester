usage：csvtomysql -config=config.toml -csv=csv.csv -log=s.log -line=500000 -sleep=3 -logtofile
-config 配置mysql数据库的信息
-csv 为导入的csv文件目录及文件名
-log 为要写入日志的路径及文件名
-line为多少行sleep一下
-sleep为sleep多少秒
-logtofile指是否将错误写入日志文件，若没有此标准则打印到屏幕
