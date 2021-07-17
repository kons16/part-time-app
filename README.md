# part-time-app
テンプレートを元にアルバイトの出退勤メッセージを生成します。  
弊研究室内のアルバイトでは、毎回メールで出退勤の連絡をする必要があり、
時間記入等々が面倒くさいので、コマンド一つで生成できるようにしました。  

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
7月17日(土) 9:00 アルバイトを開始します。
```
