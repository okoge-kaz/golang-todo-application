# Git

## git push

[git push の取り消し](https://www-creators.com/archives/2020)

### ローカル履歴から直前のコミットを取り消す

```bash
git reset --soft HEAD^
```

あまり良くないが、強制的に上書きすることができる。

```bash
git push -f origin master
```
