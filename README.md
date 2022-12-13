# Dagger.io - CNCF Turin

## init Dagger.io project

    dagger-cue project init # needed at very first time

## update project

    dagger-cue project update

## Ready to GO!

    dagger-cue do build

### detailed log

    dagger-cue do build --log-format plain

### no cache

    dagger-cue do build --log-format plain --no-cache

### subaction

    dagger-cue do build goTest
