FROM golang:1.13

RUN mkdir /goradiooooo
ADD . /goradiooooo/
WORKDIR /goradiooooo

#RUN make
