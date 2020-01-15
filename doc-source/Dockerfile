FROM ruby:2.6-alpine
EXPOSE 4567

RUN apk update \
 && apk add coreutils git make g++ nodejs

RUN gem install bundler:2.0.2

RUN git clone https://github.com/slatedocs/slate /slate/source_orig

RUN cd /slate/source_orig && bundle install

VOLUME /slate/source
VOLUME /slate/build

CMD cd /slate && cp -nr source_orig/* source && cd source && exec bundle exec middleman server --watcher-force-polling