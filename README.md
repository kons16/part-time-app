# part-time-app
テンプレートを元にアルバイトの出退勤を記録します。  
実行時引数に、開始(s) or 終了(e) と、時刻を入力して実行すると、templatesを元に文章を生成します。  
`/templates` には、`template_start.txt` と `template_end.txt` を用意して、埋め込み変数は `{}` で囲ってください。  
  
```
% go build -o reco ./main.go
% export PATH=$PATH:/path/part-time-app
% reco s 9:00
```
