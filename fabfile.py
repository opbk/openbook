import time
from fabric.api import task, local, run, sudo, settings, hide, env, prompt

def _hotfix_start():
    with hide('running', 'stdout', 'stderr'):
        local('git checkout master')
        old_version = local('cat version', capture=True)
        tmp = old_version.split('.')
        new_version = '%s.%s.%s' % (tmp[0], tmp[1], int(tmp[2]) + 1)
        new_version = prompt('Enter hotfix version:', default=new_version)

        print('Starting hotfix %s...' % new_version)
        local('git flow hotfix start ' + new_version)
        local('echo %s > version' % new_version)
        local('git commit -am \'Version switched from %s to %s\'' % (old_version, new_version))


def _hotfix_finish():
    hotfix_version = local('git flow hotfix | sed -n \'/\* /s///p\'', capture=True)
    local('git flow hotfix finish ' + hotfix_version)

@task
def hotfix(action=None):
    if action is None:
        with settings(hide('running', 'stdout', 'stderr'), ok_ret_codes=[0, 1]):
            hotfix_exist = bool(local('git flow hotfix | sed -n \'/\* /s///p\'', capture=True))
            action = 'finish' if hotfix_exist else 'start'

    if action == 'start':
        _hotfix_start()
    elif action == 'finish':
        _hotfix_finish()


def _release_start():
    with hide('running', 'stdout', 'stderr'):
        local('git checkout develop')
        old_version = local('cat version', capture=True)
        tmp = old_version.split('.')
        new_version = '%s.%s.0' % (tmp[0], int(tmp[1]) + 1)
        new_version = prompt('Enter new release version:', default=new_version)

        local('echo %s > version' % new_version)
        local('git commit -am \'Version switched from %s to %s\'' % (old_version, new_version))

        print('Starting release %s...' % new_version)
        local('git flow release start ' + new_version)
        local('git checkout release/' + new_version)


def _release_finish():
    release_version = local('git flow release | sed -n \'/\* /s///p\'', capture=True)
    local('git flow release finish ' + release_version)

@task
def release(action=None):
    if action is None:
        with settings(hide('running', 'stdout', 'stderr'), ok_ret_codes=[0, 1]):
            release_exist = bool(local('git flow release | sed -n \'/\* /s///p\'', capture=True))
            action = 'finish' if release_exist else 'start'

    if action == 'start':
        _release_start()
    if action == 'finish':
        _release_finish()