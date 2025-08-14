# 这些是该组织内的其他工具，也都是服务于kratos项目的，能够方便你的开发和使用，当然主要是我自己用
init:
	# 这个工具能够让你方便的添加proto文件
	go install github.com/orzkratos/orzkratos/cmd/orzkratos-add-proto@latest
	# 这个工具能够让你的service实现和proto接口保持同步修改，使用前请先做好代码备份或者提交代码
	go install github.com/orzkratos/orzkratos/cmd/orzkratos-srv-proto@latest
	# 这个工具能让项目升级依赖包
	go install github.com/go-mate/depbump/cmd/depbump@latest
	# 这个工具用于项目的静态检查
	go install github.com/go-mate/go-lint/cmd/go-lint@latest

all:
	cd demo1kratos && make all
	cd demo2kratos && make all

# 当你的kratos项目在proto里新增接口时，通过这个命令能够在对应的服务里也增加函数逻辑，在删除接口时也能把服务代码改为非导出的，以下是使用样例
orz:
	cd demo1kratos && orzkratos-srv-proto -auto
	cd demo2kratos && orzkratos-srv-proto -auto

# ========================================
# 同步上游仓库最新修改到 fork 项目的完整流程
# ========================================
# 背景说明：
# 1. 这些项目都是由 demokratos fork 来的，而每个fork都会展示一种特殊的使用技巧，这样方便新手查看具体如何使用。
# 2. 在新 fork 项目里，还贴心的提供了和源项目 demokratos 代码的比较的测试函数
# 3. 当源项目 demokratos 修改了东西，或者使用了更新的 kratos 框架版本时，还可以在 fork 项目里同步修改。
# 4. 因此，这些 fork 的项目，基本都不会再合并到 demokratos 里，而是作为单独的样例长期存在。
# 5. 项目提供了同步源代码修改的逻辑，这个逻辑只能在 fork 项目里执行。
# 使用说明：
# 1. 首先检查工作区状态 git status，如有未提交的修改需要处理：
#    - 仅包含 go.mod/go.sum 的变化：git stash (依赖升级可稍后再合)
#    - 有业务代码变化：需要先提交代码，避免混合提交历史
# 2. 按顺序执行 merge-step1 到 merge-step8 完成同步
# 3. 如果有冲突，自行解决 (常见于 go.mod/go.sum 文件)
# 4. 完成后提交修改：git add . && git commit --amend --no-edit (推荐，把少量依赖包升级的变动合并到merge commit)
#    或单独提交：git add . && git commit -m "简单升级依赖包"
# 5. 若任何步骤出现错误需要再次修改代码/依赖时，改完都要再次运行测试和代码静态检查，避免引入新问题

merge-step1:
	# 添加上游仓库为远程源
	# 注意: 如果 upstream 远程源已存在，而且是同名仓库，就忽略重复的错误，因为这不是问题，但是假如指向其他仓库，就报错，而且不往下执行
	git remote add upstream git@github.com:orzkratos/demokratos.git
	@echo "✅ 已添加上游仓库远程源"

merge-step2:
	# 获取上游仓库的最新代码，不获取标签以避免冲突
	git fetch --no-tags upstream main
	@echo "✅ 已获取上游仓库最新代码"

merge-step3:
	# 确保当前在 main 分支里
	git checkout main
	@echo "✅ 已切换到 main 分支里"

merge-step4:
	# 合并上游仓库的 main 分支，使用 --no-edit 避免弹出编辑器，这样适合在脚本里执行
	git merge upstream/main --no-edit
	git status
	# 【重要提醒】假如出现冲突，请严格按照以下步骤操作：
	# 1. 编辑冲突文件，逐个解决所有 <<<<<<< ======= >>>>>>> 标记的冲突
	# 【技巧策略】假如是go.sum有冲突，也可以不手动改，而是在解决完 go.mod 的冲突后执行 go mod tidy 即可解决
	# 2. 使用 git add <文件名> 将解决后的文件标记为已解决
	# 3. 继续合并流程：git merge --continue（绝对不要使用 git commit）
	# 【助手注意】在 merge 冲突场景下必须使用 git merge --continue 而非 git commit
	# 【再次强调】解决冲突后不要直接使用 git commit，这会破坏 merge 流程的状态管理
	@echo "✅ 已合并上游代码-请检查是否有冲突需要解决"

merge-step5:
	# 升级所有项目的依赖包到最新版本
	# depbump: 升级根目录依赖
	# depbump directs: 升级子项目的直接依赖
	depbump
	# 在项目根目录里进第1个项目
	cd demo1kratos && depbump directs
	# 在项目根目录里进第2个项目
	cd demo2kratos && depbump directs
	# 在升级完以后如果前面 merge 的已经完成而这里有新增变动，就把新增变动 git commit --amend 补充到 merge 的修改里
	@echo "✅ 已升级所有依赖包"

merge-step6:
	# 运行所有的测试确保代码正常工作
	# 清理测试缓存避免旧缓存影响结果
	go clean -testcache
	go test -v ./...
	# 在项目根目录里进第1个项目
	cd demo1kratos && go test -v ./...
	# 在项目根目录里进第2个项目
	cd demo2kratos && go test -v ./...
	@echo "✅ 已进行单元测试"

merge-step7:
	# 整理 go.mod 和 go.sum 文件
	# -e 参数允许在有错误时继续执行
	go mod tidy -e
	# 在项目根目录里进第1个项目
	cd demo1kratos && go mod tidy -e
	# 在项目根目录里进第2个项目
	cd demo2kratos && go mod tidy -e
	@echo "✅ 已整理所有依赖"

merge-step8:
	# 运行代码静态检查和格式化
	go-lint
	@echo "✅ 已进行代码检查"
