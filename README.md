# part-time-app
テンプレートを元にアルバイトの出退勤を記録します。  

予め、`.env` に名前等の環境変数を設定してください。  
```
NAME=田中太郎
```  
<br>

実行時引数には、開始(s) or 終了(e) と、時刻を入力して実行すると、templatesを元に文章を生成します。  
`/templates` には、`template_start.txt` と `template_end.txt` を用意して、埋め込み文字列は `{}` で囲ってください。  
  
```
% go build -o reco ./main.go
% export PATH=$PATH:/path/part-time-app

% reco s 9:00
氏名 : 田中太郎
9:00 アルバイトを開始します。
```
