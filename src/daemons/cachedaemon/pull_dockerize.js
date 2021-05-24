var gitPullCron = require('git-pull-cron');
/*
- Clone given repo into /dev/my-repo, replacing what's already there
- Schedule cron to run every weekday (Mon-Fri) at 11:30am
- When cron task runs, a `git pull origin master` will be performed
- Once cron task has run the callback will get invoked with latest commit info
 */
gitPullCron.init('git://github.com/jwilder/dockerize.git', '/var/asmweb/000000/', '00 30 11 * * 1-5', function(err, commit) {
  if (err) {
    return console.error(err.stack);
  }
  console.log('Updated to commit: ' + commit.id);
});
